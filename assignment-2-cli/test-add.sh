#!/bin/sh
i=12
while [ $i -ne 50 ]
do
        i=$(($i+1))
        journal entry add $i
done
