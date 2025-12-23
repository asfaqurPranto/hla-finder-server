# 🧬 Finding HLA

Finding HLA is a backend system designed to match organ or blood donors with patients based on HLA (Human Leukocyte Antigen) compatibility and geographical distance.  
The system is optimized using Go routines, Redis caching, and MySQL, ensuring high performance even with large datasets.

## 🚀 Features

1. **Advanced Authentication System**
   - Secure login system using email & password
   - Role-based access control (Admin / Doctor / User)

2. **Live Donor Counter**
   - Homepage shows a continuously updating donor count

3. **DNA / HLA Management**
   - Doctors/Admins can upload and manage donor HLA data
   - Each user has a unique HLA profile

4. **Advanced Matching Algorithm**
   - Matches donors with patients based on:
     - HLA compatibility (A1, A2, B1, B2, DR1, DR2)
     - Closest geographical distance
   - Uses:
     - Go routines for concurrency
     - Redis for distance caching
     - Distance API for real-world distance calculation

5. **Smart Sorting**
   - Sort match results by:
     - Best HLA match
     - Closest distance

6. **Admin Controls**
   - Admin can:
     - Update donor last donation date
     - Delete donor profiles if the donor is no longer available

## 🛠 Tech Stack

- **Language**: Go (Golang)
- **Framework**: Gin
- **Database**: MySQL
- **ORM**: GORM
- **Cache**: Redis
- **Concurrency**: Go routines & WaitGroups
- **Containerization**: Docker Compose

## API Endpoints

### Authentication & User Management

| Method | Endpoint          | Description                              |
|--------|-------------------|------------------------------------------|
| POST   | `/register`       | Register a new user                      |
| POST   | `/login`          | User login (returns authentication token)|
| GET    | `/user/info`      | Get logged-in user information           |
| PUT    | `/user/update`    | Update user profile information          |

### HLA Management & Matching

| Method | Endpoint              | Description                              |
|--------|-----------------------|------------------------------------------|
| POST   | `/hla`                | Add / update HLA data for logged-in user |
| GET    | `/hla/match`          | Get HLA match report for patient         |
| GET    | `/hla/match/sort`     | Get sorted HLA match report              |

### Admin Controls

| Method | Endpoint                           | Description                                           |
|--------|------------------------------------|-------------------------------------------------------|
| POST   | `/admin/input_hla/{user_id}`       | Admin can add/update HLA data for a user by ID        |
| POST   | `/admin/donation_date/{user_id}`   | Add or update donor last donation date (Admin only)   |
| DELETE | `/admin/delete/{user_id}`          | Delete donor profile (Admin only)                     |

### Sorting Options for `/hla/match/sort`

This endpoint accepts a json parameter for sorting ({
    "sortby":"best match or distance"
}).

| Sort Value     | Description                       |
|----------------|-----------------------------------|
| `best match`   | Sort by highest HLA match count   |
| `distance`     | Sort by nearest donor distance    |



## Project Structure
```
hla_finder/
├── cmd/app/main.go
├── internal/
│   ├── handlers/
│   ├── services/
│   ├── models/
│   ├── db/
│   ├── middleware/
├── docker/
│   └── compose.yaml
├── .env
├── README.md
```