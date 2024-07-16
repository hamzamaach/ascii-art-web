# ASCII-Art-Web

## Description

ASCII-Art-Web is a web-based implementation of the ASCII-Art project. It allows users to input text and transform it into a graphical representation using ASCII characters. The web interface supports multiple banner styles, including `shadow`, `standard`, and `thinkertoy`.

## Authors

- Hamza Maach
- ___________

## Usage

### Running the Application

1. Clone the repository:

    <!--
    ```sh
     git clone https://github.com/hamzamaach/ascii-art-web.git 
    cd ascii-art-web
    ```
    -->

2. Run the Go application:

    ```sh
    go run .
    ```

3. Open your web browser and navigate to:

    ```
    http://localhost:8080
    ```

### Web Interface

1. Enter your text in the input field.
2. Select the desired banner style (shadow, standard, or thinkertoy) using the radio buttons.
3. Click the "Submit" button to generate the ASCII art.

## Implementation Details

### Algorithm

1. **Input Parsing**: The input string is read from the user via an HTML form.
2. **Banner Selection**: The selected banner style is determined based on user input.
3. **Text Conversion**: The input string is converted to its ASCII representation using the selected banner style.
4. **Output Rendering**: The resulting ASCII art is displayed on the webpage.

### File Structure

- `main.go`: The main Go file that contains the HTTP server logic.
- `index.html`: Main page template.
- `features/`: Directory containing files that contain all functions.
- `banners/`: Directory containing banner files (`shadow.txt`, `standard.txt`, `thinkertoy.txt`).

### HTTP Endpoints

1. **GET /**: Sends the HTML response for the main page.

    - **Status Code**: 200 OK

2. **POST /ascii-art**: Receives text and banner data from the form, processes it, and returns the ASCII art.

    - **Status Code**: 
        - 200 OK: If the text is successfully converted.
        - 400 Bad Request: If the request is invalid.
        - 404 Not Found: If the banner style is not found.
        - 500 Internal Server Error: For unhandled errors.

### Error Handling

- **Invalid Input**: Returns a 400 Bad Request status.
- **Missing Banner**: Returns a 404 Not Found status.
- **Unhandled Errors**: Returns a 500 Internal Server Error status.
