# GoGator (bloGO-aggreGATOR)

Gator is an RSS feed aggregator build in Go. It is a CLI tool that allows users to: 

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a local PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of aggregated posts in the terminal, with a link to the full post

The intention of this toy project is to demonstrate an understanding of: 

- Integrating a Golang application with a PostgreSQL database.
- SQL queries and migrations. 
- Writing a long-running service that continuously makes HTTP requests and processes responses.
- Writing a long-running service that continuously retrieves and inserts data to a database.
- Using the [sqlc](https://sqlc.dev/) and [goose](https://github.com/pressly/goose) tools for creating typesafe SQL in Go.

## Getting Started (Debian Only)

### Prerequisites

You **must** have the Go toolchain installed. Instructions on how to install are [here](https://go.dev/doc/install)

### Setting up your database

1. Install PostgreSQL locally. 

```
sudo apt update
sudo apt install postgresql postgresql-contrib
```

2. Update the default password for the `postgres` system user. 

```
sudo passwd postgres
New password: 
Retype new password: 
passwd: password updated successfully
```

3. Start the postgresql service. 

```
sudo service postgresql start
```

4. Enter the `psql` shell with the `postgres` user.

```
sudo -u postgres psql
psql (16.8 (Ubuntu 16.8-0ubuntu0.24.04.1))
Type "help" for help.

postgres=#
```

5. Create a new database with the following command and connect to it. 

```
CREATE DATABASE gogator;
\c gogator
```

6. Set the user password on the database with the following command. 

```
ALTER USER postgres PASSWORD '<your-password>';
```

7. Exit the database. Install the `goose` CLI tool to set up your database tables. 

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

8. Get your connection string, as it is required by goose to interact with your database.

```
protocol://username:password@host:post/database

## Your connection string should look something like this, using the password you set in step 6.

postgres://postgres:<password>@localhost:5432/gogator
```

9. Run the `goose` UP migration to create all of the necessary tables in your PostgreSQL database.

```
cd sql/schema
goose postgres <connection-string> up

## If successful, you should see something like this:

2025/05/07 23:05:28 OK   001_users.sql (9.6ms)
2025/05/07 23:05:28 OK   002_feeds.sql (6.79ms)
2025/05/07 23:05:28 OK   003_feedfollows.sql (5.05ms)
2025/05/07 23:05:28 OK   004_last_fetched_at.sql (1.18ms)
2025/05/07 23:05:28 OK   005_posts.sql (6.54ms)
2025/05/07 23:05:28 goose: successfully migrated database to version: 5
```

### Installing the GoGator application 

You have two options when installing the GoGator application from the root directory of this project - you can either use:

- `go build` to compile a binary file in the root directory of this project, and run it with `./GoGator` 
- `go install` to allow you to use the GoGator CLI from anywhere on your machine with `GoGator`

Before you run any commands, you must create the `~/.gogatorconfig.json` file in your home directory with the following content:

```
{
    "db_url":"postgres://postgres:<password>@localhost:5432/gogator?sslmode=disable"
}
```

Once you have created this file and populated it accordingly, the first command you should run is `GoGator register <your name>` if you used `go install` or `./GoGator register <your-name>` if you used `go build` to compile a binary.  

## Command Reference

There are a number of commands you can run with GoGator. Some commands require additional arguments and some do not. Output on the terminal will tell you if this is the case or if you have entered invalid input. 

- `GoGator addfeed <feed_name> <url>` where feed_name is what you want to call the feed and url is the URL the RSS feed can be obtained.
- `GoGator agg <time_between_reqs>` - this command should be run in a separate terminal as it runs an infinite loop to refresh the RSS feeds with new content. If it is, it will scrape all of the feeds that are registered in the `feeds` table of the database at an interval of your choice. For example, running `GoGator agg 1h` will scrape all of the RSS feeds in the DB every hour. 
- `GoGator browse <limit>` - this command shows recent posts from the feeds the current user is subscribed to. The `limit` argument is optional and defines how many recent posts you want to see. It defaults to 2. 
- `GoGator feeds` - displays all feeds that have been added to the database, including name, URL and the user who added it. 
- `GoGator follow <url>` - this command sets the current user as a new follower of a feed that already exists. For example, if a different user has added a feed, you can follow that feed with this command instead of adding it again.
- `GoGator following` - this command displays the names of all of the feeds the current user is following.  
- `GoGator login <name>` - this command logs you in as a specified user. Note: there is no authentication on this application.
- `GoGator register <name>` - this command allows you to register a new user on the application and adds them to the database.  
- `GoGator unfollow <url>` - this command does the opposite to the follow command: it removes the current user as a follower of a feed. 
- `GoGator users` - this command displays all users currently registered with the application with an entry in the database.