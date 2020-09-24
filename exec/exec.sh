#!/usr/bin/bash
while [ true ]; do
    result=$(netstat -na | grep 223333 |wc -l)
    if [ $result -gt 1 ]; then
        echo "sorry,waiting for port release..."
        sleep 2
    else
      sleep 5
    ls
      echo "start serverdemo!"
      break ;
    fi
done

exit 1
