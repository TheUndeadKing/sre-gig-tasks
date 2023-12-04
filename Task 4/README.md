Following Cobra-cli is a time conversion application. It also can list the current time based on timezone.

Use Case:

(1) Download the ctime binary (Or) Clone the repo.

(2) List the all timezone by using `ctime list`.

EG:
~~~
]# ./ctime list
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

(3) To get current time based on timezone by using `ctime current TimeZone`.

EG:
~~~
]# ./ctime current Australia/Brisbane
The current time in Australia/Brisbane is 03:22
~~~

(5) Convert time & date from one zone to another zone by using `ctime convert -f FromTimezone -t ToTimezone -m HH:MM:SS -d YYYY-MM-DD`.

EG:
~~~
~]# ./ctime convert -f Europe/Amsterdam -t America/Los_Angeles -e 17:45:00 -d 2021-03-14
Time - 09:45 
Date - 03/14/2021 
Timezone - America/Los_Angeles 
~~~
