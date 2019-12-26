#/bin/bash
echo "static lib"
rm -rf ./*.o
rm -rf ./*.a
rm -rf main

gcc -c -o video_so.o video_so.c
ar -rc libvideo_so.a video_so.o
gcc main.c libvideo_so.a -L. -o static_main

rm -rf ./*.o
echo "lib"
rm -rf /usr/lib64/libvideo_so.so
gcc -fPIC -shared video_so.c -o libvideo_so.so
sudo cp libvideo_so.so /Users/wolingzong/go/src/api-service/go/
gcc main.c -L. -lvideo_so -o main
