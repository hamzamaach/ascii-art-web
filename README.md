# ASCII-Art-Web-export-file

## Description

ASCII-Art-Web-Stylize is a web-based implementation of the ASCII-Art project. It allows users to input text and transform it into a graphical representation using ASCII characters. The web interface with HTML & CSS supports multiple banner styles, including `shadow`, `standard`, and `thinkertoy`.

## Authors

- Ismail Bentour
- Hamza Maach

## Usage

### Running the Application

#### Option 1: Using Go directly

1. Clone the repository:

    ```sh
    git clone https://github.com/hamzamaach/ascii-art-web.git
    cd ascii-art-export-file
    ```

2. Run the Go application:

    ```sh
    go run .
    ```

3. Open your web browser and navigate to:

    ```
    http://localhost:8080
    ```

#### Option 2: Using Docker

1. Clone the repository:

    ```sh
    git clone https://github.com/hamzamaach/ascii-art-web.git
    cd ascii-art-export-file
    ```

2. Build the Docker image:

    ```sh
    docker build -t ascii-art-web-img .
    ```

3. Run the Docker container:

    ```sh
    docker run -d -p 8080:8080 --name ascii-art-web-con ascii-art-web-img
    ```

4. Open your web browser and navigate to:

    ```
    http://localhost:8080
    ```

To stop and remove the container:

```sh
docker stop ascii-art-web-con
docker rm ascii-art-web-con
```

### Web Interface

1. Enter your text in the input field.
2. Select the desired banner style (shadow, standard, or thinkertoy) using the dropdown selector.
3. Click the "Generate" button to generate the ASCII art or "Download" button to export the output as file.

## Implementation Details

### Algorithm

1. **Input Parsing**: The input string is read from the user via an HTML form.
2. **Banner Selection**: The selected banner style is determined based on user input.
3. **Text Conversion**: The input string is converted to its ASCII representation using the selected banner style.
4. **Output Rendering**: The resulting ASCII art can be displayed on the webpage or downloaded.

### File Structure

- `main.go`: The main Go file that contains the HTTP server logic.
- `index.html`: Main page template.
- `about.html`: About page template.
- `features/`: Directory containing files that contain all functions.
- `assets/`: Directory containing all the CSS assets.
- `banners/`: Directory containing banner files (`shadow.txt`, `standard.txt`, `thinkertoy.txt`).
- `Dockerfile`: Contains instructions for building the Docker image.

### HTTP Endpoints

1. **GET /**: Sends the HTML response for the main page.

    - **Status Code**: 200 OK

2. **POST /ascii-art**: Receives text and banner data from the form, processes it, and returns the ASCII art.

    - **Status Code**: 
        - 200 OK: If the text is successfully converted.
        - 400 Bad Request: If the request is invalid.
        - 404 Not Found: If the page or the banner style is not found.
        - 500 Internal Server Error: For unhandled errors.

### Error Handling

- **Invalid Input**: Returns a 400 Bad Request status.
- **Page not Found**: Returns a 404 Not Found status.
- **Invalid Banner**: Returns a 404 Not Found status.
- **Entity Too Large**: Returns a 413 Entity Too Large.
- **Method Not Allowed**: Returns a Method Not Allowed status.
- **Unhandled Errors**: Returns a 500 Internal Server Error status.

## Docker

The application can be containerized using Docker. The `Dockerfile` in the root directory contains the necessary instructions to build the image.

To build and run the Docker container, follow the steps in the "Usage" section under "Option 2: Using Docker".