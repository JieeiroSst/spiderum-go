#### Compile from source
2. Go to the livego directory and execute `go build` or `make build`

## Use
1. Start the service: execute the livego binary file or `make run` to start the livego service;
2. Get a channelkey(used for push the video stream) from `http://localhost:8090/control/get?room=movie` and copy data like your channelkey.
3. Upstream push: Push the video stream to `rtmp://localhost:1935/{appname}/{channelkey}` through the` RTMP` protocol(default appname is `live`), for example, use `ffmpeg -re -i demo.flv -c copy -f flv rtmp://localhost:1935/{appname}/{channelkey}` push([download demo flv](https://s3plus.meituan.net/v1/mss_7e425c4d9dcb4bb4918bbfa2779e6de1/mpack/default/demo.flv));
4. Downstream playback: The following three playback protocols are supported, and the playback address is as follows:
    - `RTMP`:`rtmp://localhost:1935/{appname}/movie`
    - `FLV`:`http://127.0.0.1:7001/{appname}/movie.flv`
    - `HLS`:`http://127.0.0.1:7002/{appname}/movie.m3u8`