# yt-transcribe

Grab a transcription of a YouTube video. Useful for when you'd like to
learn from, highlight, or take notes on a video.

## installation

Requires yt-dl and whisper-cpp.

    go install github.com/parkr/yt-transcribe/cmd/yt-transcribe

## usage

    yt-transcribe <youtube_link>

Text is written to stdout, or you can pipe to a file.
