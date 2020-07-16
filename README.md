# XFCE4 APOD Background Change Script
This script can be configured with cron to change the background of a linux XFCE4 system.

To run the script, run 

    $ ./cb_apod

This will download the APOD, save it as apod_image.jpg, and set it as XFCE4 background. 

## Configure cronjob
Follow these steps - 

    $ crontab -e
   
   Edit the file to include this line ( to run this job every day at 8 am). You can change the cron to run however you want. - 
   

    0 8 * * * /home/user/<path to cb_apod>

## Building apod binary 

If you need to rebuild the binary for some reason, get Golang and run - 

    go build 

