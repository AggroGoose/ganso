# Ganso - Go Based Backend

As I've learned more about the Go Language and the value that a backend server can provide, I decided to build this backend application for my website. The goal of this is mainly to continue learning about Go and server side programming, as my website doesn't need a separate backend for it's current needs. However if I decide to expand my website's functionality, this project with microservice architecture could also aid in that.

## Services

Starting out this project, I will have four main services that can be expanded on later:

- **Broker Service** - This service will be the entry point for the app and will determine where incoming requests need to be directed.
- **Database Service** - This service will be the core of the application and could likely be split into separate services down the line but for the time being it will handle any functionality that interacts with the DB such as managing additional post data (audio files, etc.) post events (comments, reactions, saves) and user events (user settings, etc.)
- **Ghost Service** - This service will handle interactions to/from the Ghost CMS and parsing the data. Having a separate service is useful in case I decide to switch CMS or build my own CMS.
- **Webhook Service** - The webhook service will be used to handle webhooks sent from GhostCMS when posts, tags or pages are created/updated/deleted

## Lessons Learned

To be added as I build the project.
