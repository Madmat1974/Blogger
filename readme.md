gator is a command line interface (CLI) to add users, of where the users can fetch various RSS feeds, store such feeds in a database, and post such feeds to the terminal in a user defined interval.

Dependacies required: Postgres and Go installed.

git clone https://github.com/Madmat1974/gator to clone the application.

cd into the gator folder and type go install .

Setup a gator config file by creating a file named .gatorconfig.json in your home directory. This file can be created by using a Text Editor or a command-line editor in Linux(such as Nano or Emacs). The purpose of the file is so that the gator application knows the requirements for accessing the Postgres database. Contents of the file should include a format as:
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}

To use gator, precede any commands with the word gator. Example for calling commands with description:
gator reset                       This command removes all users and feeds from the database
gator register <name>             This command adds a user to the database
gator users                       This command lists all users stored in the database with an indicator of the user that is actually logged in
gator login <name>                This command will login a registered user. Make sure to gator register <name> prior to using this command
gator agg <time interval>         This command pulls feeds that are set by the current user and displays them at the time interval to the terminal. Example gator agg 30s, gator agg 1m, gator agg 1h
gator addfeed "feedname" feedurl  Example: gator addfeed "Fox News" https://moxie.foxnews.com/google-publisher/latest.xml  The feedname is user-defined and feed is auto-followed.
gator feeds                       This lists all the feeds that the logged-in user is following by feed name and URL
gator follow <url>                Similar to gator addfeed command, but allows a user to follow a feed just by the URL and not give it a name
gator following                   Lists the feeds that are being followed by current user by feed name
gator unfollow <url>              Deletes feeds that users are following from the database
gator browse [limit]              gator browse, by itself, will default to showing two feed posts that are being followed by the user. A max of five post feeds can be displayed. Usage: gator browse 


