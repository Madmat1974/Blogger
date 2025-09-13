Dependacies required: Postgres and Go installed to run program
Go git github.com/Madmat1974/Blogger to clone the files
Within project director (using cd command), use command <go install .> without the <> tags
Setup a Blogger "Gator" config file by creating a file named .gatorconfig.json in your home directory. Ensure file is populated with the following content:

{
  "db_url": "postgres://example"
}

this will hold the url location of the Postgres database, and the file will be populated as needed for users via the application

commands avail: reset           -Removes all users and feeds from the db
                register        -adds a user useage "register $" where $ is username
                users           -lists all users in db with indicator who is logged in
                login           -logs in the user, usage "login $" $ is the username
                agg             -pull feeds by time and display, usage agg 10s (1m, 1hr)
                addfeed         -addfeed "give it a name" url
                feeds           -lists the feeds that the logged in user is polling from
                follow          -current user can follow by command follow <url>
                following       -lists the name of the feeds that are being followed
                unfollow        -removes feeds that are followed from db
                browse          -posts feeds to display, up to 5 posts with browse 5

