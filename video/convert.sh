#!/bin/bash

# 闊宠棰戝悎骞惰浆鎹㈣剼鏈紙缁熶竴鏃堕棿鎴崇増鏈級
echo "WebRTC闊宠棰戝悎骞惰浆鎹㈠伐鍏凤紙缁熶竴鏃堕棿鎴筹級"
echo "====================================="

# 妫€鏌ユ槸鍚︽湁IVF鏂囦欢
ivf_files=$(ls file/*/*video.ivf 2>/dev/null)
if [ -z "$ivf_files" ]; then
    echo "娌℃湁鎵惧埌video.ivf鏂囦欢"
    exit 1
fi

echo "鎵惧埌浠ヤ笅瑙嗛鏂囦欢锛�"
echo "$ivf_files"
echo ""

# 澶勭悊姣忎釜IVF鏂囦欢
for ivf_file in $ivf_files; do
    echo "澶勭悊鏂囦欢: $ivf_file"

    # 鐢熸垚鏂囦欢鍚�
    base_name="${ivf_file%video.ivf}"
    audio_ogg="${base_name}audio.ogg"
    audio_rtp="${base_name}audio.rtp"
    video_rtp="${base_name}video.rtp"
    webm_file="${base_name}video.webm"
    mp4_file="${base_name}video.mp4"
    combined_webm="${base_name}combined.webm"
    combined_mp4="${base_name}combined.mp4"

    echo "瑙嗛鏂囦欢: $ivf_file"
    echo "闊抽OGG鏂囦欢: $audio_ogg"
    echo "闊抽RTP鏂囦欢: $audio_rtp"
    echo "瑙嗛RTP鏂囦欢: $video_rtp"

    # 妫€鏌ユ槸鍚︽湁瀵瑰簲鐨勯煶棰戞枃浠�
    if [ -f "$audio_ogg" ] || [ -f "$audio_rtp" ]; then
        if [ -f "$audio_ogg" ]; then
            echo "鉁� 鎵惧埌OGG闊抽鏂囦欢: $audio_ogg"
            audio_file="$audio_ogg"
        else
            echo "鉁� 鎵惧埌RTP闊抽鏂囦欢: $audio_rtp"
            audio_file="$audio_rtp"
        fi

        # 鍏堣浆鎹㈣棰�
        echo "杞崲瑙嗛涓篧ebM..."
        ffmpeg -i "$ivf_file" -c:v copy "$webm_file" 2>/dev/null

        if [ $? -eq 0 ]; then
            echo "鉁� 瑙嗛杞崲鎴愬姛: $webm_file"

            # 灏濊瘯鍚堝苟闊宠棰�
            echo "鍚堝苟闊宠棰戜负WebM..."
            ffmpeg -i "$webm_file" -i "$audio_file" -c:v copy -c:a copy "$combined_webm" 2>/dev/null

            if [ $? -eq 0 ]; then
                echo "鉁� 闊宠棰戝悎骞舵垚鍔�: $combined_webm"
            else
                echo "鉂� 闊宠棰戝悎骞跺け璐ワ紝灏濊瘯閲嶆柊缂栫爜..."

                # 灏濊瘯閲嶆柊缂栫爜闊抽
                ffmpeg -i "$webm_file" -i "$audio_file" -c:v copy -c:a aac "$combined_mp4" 2>/dev/null
                if [ $? -eq 0 ]; then
                    echo "鉁� 闊宠棰戝悎骞朵负MP4鎴愬姛: $combined_mp4"
                else
                    echo "鉂� 闊宠棰戝悎骞跺け璐ワ紝浣嗚棰戞枃浠跺彲鐢�"
                fi
            fi
        else
            echo "鉂� 瑙嗛杞崲澶辫触"
        fi

        # 鍗曠嫭杞崲闊抽锛堢敤浜庤皟璇曪級
        echo "杞崲闊抽涓篧AV..."
        audio_wav="${base_name}audio.wav"
        ffmpeg -i "$audio_file" -c:a pcm_s16le "$audio_wav" 2>/dev/null
        if [ $? -eq 0 ]; then
            echo "鉁� 闊抽杞崲鎴愬姛: $audio_wav"
        else
            echo "鉂� 闊抽杞崲澶辫触"
        fi

    else
        echo "鉂� 娌℃湁鎵惧埌瀵瑰簲鐨勯煶棰戞枃浠�"
        echo "鍙浆鎹㈣棰戞枃浠�..."

        # 鍙浆鎹㈣棰�
        ffmpeg -i "$ivf_file" -c:v copy "$webm_file" 2>/dev/null
        if [ $? -eq 0 ]; then
            echo "鉁� 瑙嗛杞崲鎴愬姛: $webm_file"
        else
            echo "鉂� 瑙嗛杞崲澶辫触"
        fi
    fi

    echo ""
done

echo "杞崲瀹屾垚锛�"
echo ""
echo "鏂囦欢璇存槑锛�"
echo "- *video.webm: 绾棰戞枃浠讹紙VP8缂栫爜锛�"
echo "- *audio.ogg: 绾煶棰戞枃浠讹紙OGG缂栫爜锛�"
echo "- *combined.webm: 闊宠棰戝悎骞舵枃浠讹紙鎺ㄨ崘锛�"
echo "- *combined.mp4: 闊宠棰戝悎骞舵枃浠讹紙MP4鏍煎紡锛�"
echo "- *audio.wav: 闊抽璋冭瘯鏂囦欢"
echo ""
echo "鎾斁寤鸿锛�"
echo "1. 浼樺厛浣跨敤 combined.webm 鏂囦欢锛堝寘鍚煶瑙嗛锛�"
echo "2. 濡傛灉娌℃湁闊抽锛屼娇鐢ㄧ函瑙嗛鐨� .webm 鏂囦欢"
echo "3. 浣跨敤VLC鎾斁鍣ㄦ挱鏀炬墍鏈夋牸寮�"
echo ""
echo "浼樺娍锛�"
echo "鉁� 缁熶竴鏃堕棿鎴崇‘淇濋煶瑙嗛鏂囦欢鍖归厤"
echo "鉁� 鑷姩鍚屾闊宠棰戣建閬�"
echo "鉁� 鏀寔澶氱杈撳嚭鏍煎紡"