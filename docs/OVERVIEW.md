# Project Overview

This project provides a minimalist, client-side web application designed for quick, on-demand generation of Universally Unique Identifiers (UUIDs). It serves as a straightforward tool for developers or anyone needing a UUID without relying on external services or complex local setups.

## Architecture

The application is a single-page web utility implemented entirely within `index.html`. It operates as a static file, executing all its logic directly within the user's web browser. There is no backend server component, database, or complex build process involved.

## Key Files

*   `index.html`: This is the sole application file. It contains all the necessary HTML structure for the user interface, any embedded CSS for styling, and the JavaScript logic responsible for generating and displaying UUIDs.
*   `README.md`: Provides a concise overview of the project's purpose and includes a screenshot of the application.
*   `CNAME`: Specifies a custom domain name for hosting the application, typically used in conjunction with GitHub Pages.

## How to Run

To run this application locally, follow these steps:

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/calvinbrown085/uuid-gen.git
    ```
2.  **Navigate into the project directory:**
    ```bash
    cd uuid-gen
    ```
3.  **Open in a browser:**
    Simply open the `index.html` file directly in your preferred web browser.
    Alternatively, for a more standard local serving environment, you can use a simple static file server:
    *   **Using Python (if installed):**
        ```bash
        python -m http.server 8000
        ```
        Then, navigate to `http://localhost:8000` in your browser.
    *   **Using Node.js `serve` (if installed globally):**
        ```bash
        npx serve .
        ```
        Then, navigate to the URL provided by `serve`.

## How to Test

No automated test suite is provided with this repository. The application can be manually tested by:

1.  Opening `index.html` in a web browser.
2.  Interacting with the UI element designed to generate UUIDs (e.g., a button).
3.  Verifying that a new, correctly formatted UUID is displayed upon each generation.