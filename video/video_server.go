package video

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media/ivfwriter"
	"github.com/pion/webrtc/v4/pkg/media/oggwriter"
)

func RegisterVideoOffer(r *gin.Engine) {
	log.Println("RegisterVideo..")
	r.StaticFile("/index", "./video/static/index.html")
	r.StaticFile("/favicon.ico", "./video/static/favicon.ico")
	//http.Handle("/index", http.FileServer(http.Dir("./static")))
	r.POST("/offer", func(context *gin.Context) {
		handleOffer(context.Writer, context.Request)
	})

	//r.POST("/offer", gin.WrapF(handleOffer))

	r.POST("/offer/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "test request received"})
		return
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("./video/static/index.html")
	})
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/offer", handleOffer)

	addr := ":8080"
	log.Printf("Server listening at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleOffer(w http.ResponseWriter, r *http.Request) {
	var offer webrtc.SessionDescription
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		http.Error(w, "invalid offer", http.StatusBadRequest)
		return
	}

	//如果 SDP 为空
	if len(offer.SDP) == 0 {
		http.Error(w, "SDP cannot be empty", http.StatusBadRequest)
		return
	}

	// 创建 PeerConnection
	pc, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{URLs: []string{"stun:stun.l.google.com:19302"}},
		},
	})
	if err != nil {
		panic(err)
	}

	// 创建统一的会话时间戳
	sessionTime := strconv.FormatInt(time.Now().UnixMilli(), 10)
	fmt.Printf("Starting new recording session: %s\n", sessionTime)

	// 创建文件路径
	videoRtpPath := filepath.Join(".", "video", "file", sessionTime, "video.rtp")
	videoIvfPath := filepath.Join(".", "video", "file", sessionTime, "video.ivf")
	audioOggPath := filepath.Join(".", "video", "file", sessionTime, "audio.ogg")
	audioRtpPath := filepath.Join(".", "video", "file", sessionTime, "audio.rtp")

	// 先删除可能存在的旧文件
	_ = os.Remove(videoRtpPath)
	_ = os.Remove(videoIvfPath)
	_ = os.Remove(audioOggPath)
	_ = os.Remove(audioRtpPath)

	// 创建writers
	var videoIvfWriter *ivfwriter.IVFWriter
	var audioOggWriter *oggwriter.OggWriter
	var videoRtpFile, audioRtpFile *os.File

	dir := filepath.Dir(videoRtpPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		panic(err)
	}

	// 创建视频文件
	videoRtpFile, err = os.Create(videoRtpPath)
	if err != nil {
		fmt.Printf("Failed to create video RTP file: %v\n", err)
		return
	}

	videoIvfWriter, err = ivfwriter.New(videoIvfPath, ivfwriter.WithCodec("video/VP8"))
	if err != nil {
		fmt.Printf("Failed to create IVF writer: %v\n", err)
		videoIvfWriter = nil
	}

	// 创建音频文件
	audioRtpFile, err = os.Create(audioRtpPath)
	if err != nil {
		fmt.Printf("Failed to create audio RTP file: %v\n", err)
		return
	}

	oggFile, err := os.Create(audioOggPath)
	if err != nil {
		fmt.Printf("Failed to create OGG file: %v\n", err)
		return
	}

	audioOggWriter, err = oggwriter.NewWith(oggFile, 48000, 2)
	if err != nil {
		fmt.Printf("Failed to create OGG writer: %v\n", err)
		audioOggWriter = nil
	}

	fmt.Printf("Files created successfully:\n")
	fmt.Printf("  Video RTP: %s\n", videoRtpPath)
	fmt.Printf("  Video IVF: %s\n", videoIvfPath)
	fmt.Printf("  Audio OGG: %s\n", audioOggPath)
	fmt.Printf("  Audio RTP: %s\n", audioRtpPath)

	// 收到远端 track
	pc.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		fmt.Printf("Got remote track: %s (kind: %s)\n", track.Codec().MimeType, track.Kind())

		go func() {
			// 在goroutine内部管理文件关闭
			defer func() {
				if track.Kind() == webrtc.RTPCodecTypeVideo {
					_ = videoRtpFile.Close()
					if videoIvfWriter != nil {
						_ = videoIvfWriter.Close()
					}
				} else if track.Kind() == webrtc.RTPCodecTypeAudio {
					_ = audioRtpFile.Close()
					if audioOggWriter != nil {
						_ = audioOggWriter.Close()
					}
					_ = oggFile.Close()
				}
			}()

			fmt.Printf("Starting to read %s packets...\n", track.Kind())
			packetCount := 0
			for {
				rtp, _, err := track.ReadRTP()
				if err != nil {
					fmt.Printf("%s ReadRTP error: %v\n", track.Kind(), err)
					return
				}
				packetCount++
				if packetCount%100 == 0 {
					fmt.Printf("Processed %d %s packets\n", packetCount, track.Kind())
				}

				if track.Kind() == webrtc.RTPCodecTypeVideo {
					// 写入视频RTP数据
					rtpData, err := rtp.Marshal()
					if err != nil {
						fmt.Printf("Video Marshal RTP error: %v\n", err)
						continue
					}

					length := uint32(len(rtpData))
					if err := binary.Write(videoRtpFile, binary.BigEndian, length); err != nil {
						fmt.Printf("Video Write length error: %v\n", err)
						return
					}

					if _, err := videoRtpFile.Write(rtpData); err != nil {
						fmt.Printf("Video Write RTP file error: %v\n", err)
						return
					}

					// 写入IVF文件
					if videoIvfWriter != nil {
						if err := videoIvfWriter.WriteRTP(rtp); err != nil {
							fmt.Printf("IVF WriteRTP error: %v\n", err)
						}
					}

				} else if track.Kind() == webrtc.RTPCodecTypeAudio {
					// 写入音频RTP数据
					rtpData, err := rtp.Marshal()
					if err != nil {
						fmt.Printf("Audio Marshal RTP error: %v\n", err)
						continue
					}

					length := uint32(len(rtpData))
					if err := binary.Write(audioRtpFile, binary.BigEndian, length); err != nil {
						fmt.Printf("Audio Write length error: %v\n", err)
						return
					}

					if _, err := audioRtpFile.Write(rtpData); err != nil {
						fmt.Printf("Audio Write RTP file error: %v\n", err)
						return
					}

					// 写入OGG文件
					if audioOggWriter != nil {
						if err := audioOggWriter.WriteRTP(rtp); err != nil {
							fmt.Printf("OGG WriteRTP error: %v\n", err)
						}
					}
				}
			}
		}()
	})

	// 设置远端描述
	fmt.Println("SetRemoteDescription offer")
	if err := pc.SetRemoteDescription(offer); err != nil {
		panic(err)
	}

	// 创建 answer
	fmt.Println("CreateAnswer")
	answer, err := pc.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	pc.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		fmt.Println("PeerConnection state:", state)
		if state == webrtc.PeerConnectionStateDisconnected ||
			state == webrtc.PeerConnectionStateFailed ||
			state == webrtc.PeerConnectionStateClosed {
			//这里处理mp4转换

			fmt.Println("❌ 连接已断开，所有轨道可能已停止")
		}
	})

	// 等待 ICE 收集完成
	fmt.Println("GatheringCompletePromise")
	gatherComplete := webrtc.GatheringCompletePromise(pc)
	if err := pc.SetLocalDescription(answer); err != nil {
		panic(err)
	}
	<-gatherComplete

	fmt.Println("GatherComplete and resp answer")
	resp, _ := json.Marshal(pc.LocalDescription())
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
