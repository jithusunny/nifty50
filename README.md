A simple Web app that displays the Top 10 Gainers and Losers from [here](http://nseindia.com/live_market/dynaContent/live_analysis/top_gainers_losers.htm)

It reads this data from a DB which would be populated by a background running process.



* /bin/fetchnifty50.go is a go program that can be run in the background (maybe every 5 minutes) to fetch data from the NSEIndia website and push it into the DB.

* fetchnifty50.go may be crontab-ed in single machine deployment or executed in this way - [Rethinking Cron](http://adam.heroku.com/past/2010/4/13/rethinking_cron/) - in horizontally scalable environments.

* This Go app is created using the Beego framework.
