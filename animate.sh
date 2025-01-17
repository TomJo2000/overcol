#!/bin/env bash

input_folder="images"
output_file="images/slideshow.mp4"
framerate=${1:-30}
bitrate="6000k"
crf="1" #constant rate factor
preset="veryslow"


ffmpeg -framerate $framerate -i $input_folder/%d.png -vf "pad=ceil(iw/2)*2:ceil(ih/2)*2,boxblur=10:1,format=yuv420p" -r $framerate -b:v $bitrate -crf $crf -preset $preset $output_file
