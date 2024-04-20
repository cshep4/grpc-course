const {FileUploadServiceClient} = require('./proto/streaming_grpc_web_pb');
const {UploadFileRequest, DownloadFileRequest} = require('./proto/streaming_pb');

const client = new FileUploadServiceClient('http://localhost:50051');

function uploadFile() {
    const fileInput = document.getElementById('fileUpload');
    const file = fileInput.files[0];
    if (!file) {
        alert("Please select a file.");
        return;
    }

    const stream = client.uploadFile({}, {});
    const reader = new FileReader();
    reader.onload = (e) => {
        const chunk = new UploadFileRequest();
        chunk.setContent(new Uint8Array(e.target.result));
        stream.write(chunk);
    };
    reader.onloadend = () => {
        stream.end();
    };
    reader.readAsArrayBuffer(file);
}

function downloadFile() {
    const filename = document.getElementById('fileDownload').value;
    if (!filename) {
        alert("Please enter a filename.");
        return;
    }

    const req = new DownloadFileRequest();
    req.setName(filename);
    const stream = client.downloadFile(req, {});
    stream.on('data', function (response) {
        // Handle data - for example, save it to a file or display it
        console.log(response.getContent_asU8());
    });
}

// Add error handling and UI updates as needed
