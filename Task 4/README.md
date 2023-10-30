Following go program is build using crobra-cli. It can list the current time based on timezone & convert date & time from one timezone to another timezone

(1) Clone the repo:

~~~
~]# https://github.com/TheUndeadKing/sre-gig-tasks.git
~~~

(2) Switch into `Task 4` directory.

~~~
~]# cd Task 3
~~~

(3) List the all timezone by using `go run main.go list`.

EG:
~~~
]# go run main.go list
Africa/Abidjan
Africa/Algiers
Africa/Bissau
Africa/Cairo
[..]
Pacific/Palau
Pacific/Pitcairn
Pacific/Port_Moresby
Pacific/Rarotonga
Pacific/Tahiti
Pacific/Tarawa
Pacific/Tongatapu
~~~

(4) To get current time based on timezone by using `go run main.go currentTime`.

EG:
~~~
]# go run main.go currentTime
Enter Timezone: Australia/Brisbane
The current time in Australia/Brisbane is 17:00
~~~

(5) Convert time & date from one zone to another zone by using `go run main.go convert`.

EG:
~~~
~]# go run main.go convert
Enter from timezone[EG: Europe/Amsterdam]:Australia/Brisbane
Enter from date[FORMAT: YYYY-MM-DD][EG: 2021-03-14]:2023-10-30
Enter from time[FORMAT: HH:MM:SS][EG: 17:45:00]:12:33:00
Enter to timezone[EG: America/Los_Angeles]:Asia/Kolkata
Time - 08:03 
Date - 10/30/2023 
Timezone - Asia/Kolkata 
~~~
