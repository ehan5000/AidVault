# AidVault
## 📦 What is AidVault?

**AidVault** is a secure backend system (API) that helps organizations manage and track humanitarian aid distribution — like food, shelter, or medicine — to people in need.  
It’s designed to make the process of requesting, managing, and verifying aid fast, reliable, and trackable.

---

## 👥 Who is it for?

### 1. People in need (recipients)
- Submit aid requests (e.g., food, shelter, medicine)
- Optionally upload proof documents (like ID or medical note)
- Get notified when their request is received and later fulfilled

### 2. Organizations (NGOs, charities, food banks)
- Register themselves in the system
- View submitted aid requests assigned to them
- Mark aid as fulfilled, track aid history
- Export and monitor the status of aid delivery

---

## 🌍 Example Use Case

Let’s say:
- You're someone displaced after a flood in your city.
- You go to a partner NGO website and fill a form: “Need shelter and food”
- Behind the scenes, that form hits AidVault’s `/aid-request` API
- The NGO sees your request in their dashboard (powered by AidVault)
- Once they provide aid, they mark the request as **fulfilled**
- You check the status using `/aid-status` or `/aid-request/:id/status`

---


## 🛠️ How to Run the Project

### 📦 Requirements
- [Golang](https://go.dev/doc/install) (v1.18+ recommended)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) with Docker Compose
- Internet connection (to pull PostgreSQL Docker image)

---

### 🧪 Steps to Run

1. **Clone the repository**

```bash
git clone https://github.com/your-username/aidvault.git
cd aidvault
```

2. **Start PostgreSQL using Docker Compose**

```bash
docker-compose up
```

3. **Run the Go backend**

Open a second terminal window and run:

```bash
go run main.go
```

✅ You should see:
```
Connected to PostgreSQL
✅ AidVault API running on port 8080
```

---

### 🌐 Test the API (Optional)

- Visit [https://editor.swagger.io](https://editor.swagger.io)
- Paste the contents of `docs/swagger.yaml`
- Use “Try it out” to test all endpoints interactively

---

## ⚙️ Technical Overview

- REST API in **Golang**
- Stores data in **PostgreSQL**
- Uses **AWS S3** for file uploads (e.g. ID verification)
- **Dockerized** for easy setup
- Swagger/OpenAPI docs for testing
- GitHub Actions CI/CD setup
- ✅ Ready for cloud deployment (e.g., **AWS**, **Render**, **Railway**)

## 📜 License

This project is released under a custom license.  
You are free to use, modify, and share it **as long as you credit the original author: _Ehan Hassan_**.  
Commercial use or redistribution without permission is not allowed.