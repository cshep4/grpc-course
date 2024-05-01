import './App.css'
import {createGrpcWebTransport} from "@connectrpc/connect-web";
import {createPromiseClient} from "@connectrpc/connect";
import {FileUploadService} from "./proto/streaming_connect.ts";
import {DownloadFileRequest, UploadFileRequest} from "./proto/streaming_pb.ts";
import {useState} from "react";

const transport = createGrpcWebTransport({
    baseUrl: "http://localhost:50051",
    useBinaryFormat: true,
});

const client = createPromiseClient(FileUploadService, transport);

function App() {
    const [file, setFile] = useState<File>()
    const [fileName, setFileName] = useState<string>()

    function handleChange(event: any) {
        setFile(event.target.files[0])
    }

    function handleFileNameChange(event: any) {
        setFileName(event.target.value)
    }

    async function uploadFile() {
        if (!file) {
            console.error("file must be provided")
            return
        }

        const res = await client.uploadFile(createUploadStream(file));
        console.log(res.name)
    }

    async function downloadFile() {
        if (!fileName) {
            console.error("file name must be provided")
            return
        }

        const res = client.downloadFile(new DownloadFileRequest({name: fileName}));

        // Array to hold file chunks
        const chunks: Uint8Array[] = [];
        try {
            // Iterate over the stream to receive chunks
            for await (const next of res) {
                // Collect chunks into an array
                chunks.push(next.content as Uint8Array);
            }

            // Combine all chunks into a single Blob
            const blob = new Blob(chunks, {type: 'application/octet-stream'});

            // Create a URL for the Blob
            const url = window.URL.createObjectURL(blob);

            // Create a temporary link element and trigger the download
            const a = document.createElement('a');
            a.href = url;
            a.download = fileName;  // Set the default file name for saving
            document.body.appendChild(a);  // Append the link to the document
            a.click();  // Trigger the download
            document.body.removeChild(a);  // Clean up the document

            // Revoke the Blob URL to free up resources
            window.URL.revokeObjectURL(url);
        } catch (error) {
            console.error('Failed to download file:', error);
        }
    }


    return (
        <>
            <div className="container mt-5">
                <h1 className="mb-4">File Upload Application</h1>

                <div className="mb-3">
                    <h2 className="form-label">Upload File</h2>
                    <input className="form-control" type="file" onChange={handleChange} id="fileUpload"/>
                    <button className="btn btn-primary mt-2" onClick={uploadFile}>Upload</button>
                </div>

                <div>
                    <h2 className="form-label">Download File</h2>
                    <input className="form-control" type="text" id="fileDownload" onChange={handleFileNameChange}
                           placeholder="Enter filename"/>
                    <button className="btn btn-primary mt-2" onClick={downloadFile}>Download</button>
                </div>
            </div>
        </>
    )
}

async function* createUploadStream(file: File): AsyncIterable<UploadFileRequest> {
    const chunkSize = 1024 * 64; // 64KB chunk size
    let position = 0;

    while (position < file.size) {
        const slice = file.slice(position, position + chunkSize);
        const buffer = await slice.arrayBuffer();
        const request = new UploadFileRequest({
            content: new Uint8Array(buffer)
        });
        yield request;
        position += chunkSize;
    }
}

export default App
