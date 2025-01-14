# Personal Portfolio Website  

A modern and customizable portfolio website designed to showcase personal projects, blog posts, and skills. Built with Go and Echo for backend efficiency, templ for modular server-side rendering, and HTMX for dynamic client-side updates.  

## Features  

- **Dynamic Portfolio**: Display projects with descriptions, links, and categorization.  
- **Custom Blog Engine**: Manage and display Markdown-based blog posts with sorting and individual pages.  
- **Interactive Components**: Modular UI built with templ, offering reusable and responsive layouts.  
- **Static and Dynamic Content**: Serve static assets (e.g., CV, icons) and dynamically rendered pages.  
- **Cross-Origin Support**: Configured CORS to support modern web development practices.  
- **Hot Reloading**: Streamlined development with Golang Air.  

## Technologies  

### Backend  
- **Go**: Efficient server-side logic and routing.  
- **Echo**: High-performance web framework for HTTP handling.  
- **templ**: Component-based server-side rendering engine.  
- **Golang Air**: Enables live-reloading for development.  

### Frontend  
- **HTMX**: Handles dynamic content updates on the client side.  
- **CSS**: Custom styles for responsive and user-friendly design.  

### Additional Tools  
- **Markdown Processing**: Goldmark library for rendering Markdown content.  
- **Docker**: Simplifies deployment and ensures consistent performance.  

## Project Structure  

```
project_root/  
├── cmd/website_v4/         # Entry point for the application  
├── data/                   # Blog posts and other data files  
│   ├── blog/               # Markdown blog posts  
│   └── projects/           # Project data  
├── internal/               # Application logic  
│   ├── data/               # Skills, projects, and contacts data  
│   ├── helper/             # Utility functions  
│   ├── models/             # Data models  
│   └── routes/             # Routing logic for Echo  
├── web/                    # UI components and static assets  
│   ├── components/         # templ components for UI  
│   ├── public/             # Static files (CSS, assets)  
│   └── css/                # Stylesheets  
├── .air.toml               # Configuration for Golang Air  
├── .gitignore  
├── Dockerfile              # Containerization configuration  
├── go.mod                  # Module dependencies  
└── go.sum                  # Module checksums  
```

## Getting Started  

### Prerequisites  
- **Go**: Version 1.23 or later.  
- **Docker**: For containerized deployment (optional).  
- **Golang Air**: For hot reloading during development.  

### Setup  

1. Clone the repository:  
   ```bash  
   git clone https://github.com/deej-tsn/website_v4.git  
   cd website_v4  
   ```  

2. Install dependencies:  
   ```bash  
   go mod download  
   ```  

3. Run the application with hot reloading:  
   ```bash  
   air  
   ```  

   The application runs by default at `http://localhost:8080`.  

### Using Docker  

1. Build and run the application using Docker:  
   ```bash  
   docker build -t personal-portfolio .  
   docker run -p 8080:8080 personal-portfolio  
   ```  

## Contributing  

Contributions are welcome! Fork the repository, make changes, and open a pull request.  

## License  

This project is licensed under the [MIT License](LICENSE).  

## Acknowledgements  

- [Echo](https://echo.labstack.com/) - Web framework for Go  
- [templ](https://github.com/a-h/templ) - Component-based rendering engine  
- [HTMX](https://htmx.org/) - Lightweight library for client-side interactivity  
- [Goldmark](https://github.com/yuin/goldmark) - Markdown processor  
- [Golang Air](https://github.com/cosmtrek/air) - Live-reloading for Go applications  

---