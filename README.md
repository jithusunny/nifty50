A simple Web app that displays the Top 10 Gainers and Losers from [here](http://nseindia.com/live_market/dynaContent/live_analysis/top_gainers_losers.htm)

It reads this data from a DB which would be populated by a background running process.



#### Note about the Background job program - fetchnifty50.go
/worker/fetchnifty50.go is a go program that can be run in the background (maybe every 5 minutes) to fetch data from the NSEIndia website and push it into the DB.

Here the fetchnifty50.go code has been merged into the controllers/deafult.go file.
There it can be seen that the code from fetchnifty50.go is run as part of the '/' route handler.

This hack so that the app can run on Heroku until I sort out the Heroku account settings to enable the scheduler.



* The /worker/fetchnifty50.go may be crontab-ed in single machine deployment or executed in this way - [Rethinking Cron](http://adam.heroku.com/past/2010/4/13/rethinking_cron/) - in horizontally scalable environments.

* This Go app is created using the Beego framework.
