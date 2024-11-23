# **Parallelizing MongoDB to Postgres Data Transfer**

This application is designed to transfer data from a **MongoDB** collection to a **Postgres** database efficiently. The application leverages concurrency in Go to handle large datasets by processing data in parallel batches.

---

## **Features**
- High-performance data transfer from **MongoDB** to **Postgres**.
- Optimized for large datasets with support for batching and parallel processing.
- Automatically creates the target table in Postgres if it doesn’t exist.
- Fully containerized setup using **Docker Compose**.

---

## **Technologies Used**
- **Go**: Core language for the application.
- **MongoDB**: Source database for the data transfer.
- **Postgres**: Target database relational sql.
- **Docker**: Containerization for MongoDB and ClickHouse.
- **Makefile**: Automates setup and execution tasks.

---

## **Application Workflow**
1. Connects to a MongoDB database and counts the documents in a collection.
2. Reads documents in batches from MongoDB.
3. Transfers the data to Postgres using bulk inserts.
4. Utilizes multiple workers to parallelize the data transfer, ensuring high performance.

---

## **Setup Instructions**

### **Prerequisites**
1. Install **Go** (1.18 or higher).
2. Install **Docker** and **Docker Compose**.
3. Clone this repository:
   ```bash
   $ git clone https://github.com/G-Villarinho/parallelizing-golang-ops.git
   $ cd parallelizing-golang-ops


## **How to run**

    $ make start

   
