#!/bin/sh

while read line; do

echo "drone build promote $DRONE_REPO $DRONE_BUILD_NUMBER $ENVIRONMENT --param=tag=$line"
drone build promote $DRONE_REPO $DRONE_BUILD_NUMBER $ENVIRONMENT --param=TAG=$line

done < $FILE