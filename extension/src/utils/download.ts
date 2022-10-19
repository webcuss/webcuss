export const downloadCredential = (username: string, password: string) => {
    const blob = new Blob(
        [
            "Below is your Webcuss credential, keep it safe! :)",
            "\n",
            "\n",
            "username: " + username,
            "\n",
            "password: " + password
        ],
        {
            type: "text/plain"
        }
    );
    const url = URL.createObjectURL(blob);
    chrome.downloads.download({
        url: url,
        filename: "webcuss-credential.txt"
    });
}
