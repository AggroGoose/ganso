# Ganso - Go Based Backend

As I've learned more about the Go Language and the value that a backend server can provide, I decided to build this backend application for my website. The goal of this is mainly to continue learning about Go and server side programming, as my website doesn't need a separate backend for it's current needs. However if I decide to expand my website's functionality, this project with microservice architecture could also aid in that.

## Services

Initially this project would have been broken into 4 services (5 including the DB) but after some consideration I decided to stick with one core service where the bulk of the app would be built, but leave the door open for additional services or splitting into multiple services at a later date. This core app will handle interface with external requests, DB Processing, and webhooks.

## Lessons Learned

- While learning about gRPC and other ways of making microservices communicate together, I realized at current scale of the project that a microservices architecture might be less performant than a more centralized single app. I will leave the door open for microservices as the application grows but at present time it may have more pitfalls than not.
- Also learned that before I can implement integration with GhostCMS I will need more proficiency with JWT WebTokens as that's the preferred connection method with the GhostCMS Admin API. For the meantime I will keep that application within the Next.JS Application running the website.
- Found that I favored the 'OOP' model for the Server and DB Queries that I learned in the 'Simple-Bank' project. Refactored router from **[Go-Chi](https://github.com/go-chi/chi)** to **[Gin](https://github.com/gin-gonic/gin)** to facilitate this. This model also includes a query pool for backing out of long transactions such as deleting a post or user with multiple queries.
- Following up that last statement, I learned I don't have to do a transfer pool for deletes. By adding the 'ON DELETE CASCADE' option to the foreign key, replies, comments, etc. will automatically delete when one of their dependencies is deleted.
