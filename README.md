# cc-library-management-system

- This project aims to implement a book rental application (hereinafter referred to as the library), containerize it, and deploy it to a local Kubernetes cluster using Minikube.


## Table of contents
* [Functionality Overview ](#functionality-overview)
* [Infrastructure Setup](#infrastructure-setup)
* [Architecture](#architecture)

### Functionality Overview 
- The application consists of the following components:
1. Belgrade library,
2. Novi Sad library,
3. Niš library and
4. Central library.

 1.Member Registration:
 - After entering the required information into the form (name, surname, address, and ID number), a request is sent to the central library application to check if a member with the given ID number is already registered in the central library database.
 - If the member does not exist in the central library database, they are registered, and a successful registration response is sent to the requesting city library application.
 - If the member already exists in the central library database, they are not registered, and an unsuccessful registration response is sent to the requesting city library application.
 2. Book Rental:
 - After entering the required information into the form (book title, author, ISBN, rental date, and member ID), a request is sent to the central library application to check if the member with the given ID has already borrowed the maximum of 3 books.
 - If the member has borrowed less than 3 books, the information about the new loan is recorded in the central library database, and a response is sent to the requesting city library application. 
 - Additionally, the information about the book rental by the library member is stored in the city library database.
 - If the member has borrowed the maximum of 3 books, an unsuccessful loan response is sent to the requesting city library application.
    * The central library database stores information about registered members
    * The city library databases store information about specific book rentals by specific members

### Infrastructure Setup
Before running this script, make sure you have Docker Compose, Minikube, and kubectl installed.

Executing ./run.sh will perform the following tasks:

1. Initialize a Minikube cluster,
2. Deploy resources for the Central, Niš, Belgrade, and Novi Sad libraries,
3. Apply Ingress configuration to manage external traffic,
4. Display the host information for the Ingress

### Architecture
![architecture](/docs/arch.png)
