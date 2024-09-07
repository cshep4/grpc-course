'use client';

import {CSSProperties, FormEvent, useState} from 'react';
import {useClient} from '@/grpc/client';
import {HelloService} from "@/proto/hello_connect";

const Home = () => {
    const [name, setName] = useState('');
    const [message, setMessage] = useState('');
    const client = useClient(HelloService)

    const handleSubmit = async (event: FormEvent) => {
        event.preventDefault();
        try {
            const res = await client.sayHello({name: name});
            setMessage(res.message);
        } catch (error) {
            console.error(error);
            setMessage('Error: ' + (error as Error).message);
        }
    };

    return (
        <div style={formWrapper}>
            <h1>gRPC-Web Client</h1>
            <form onSubmit={handleSubmit} style={formStyle}>
                <input
                    type="text"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    placeholder="Enter your name"
                    style={inputStyle}
                />
                <button type="submit" style={buttonStyle}>Say Hello</button>
            </form>
            {message && <p>Message from server: {message}</p>}
        </div>
    );
};

const formWrapper: CSSProperties = {
    padding: '20px',
    textAlign: 'center',
}

const formStyle: CSSProperties = {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
};

const inputStyle: CSSProperties = {
    padding: '10px',
    margin: '10px 0',
    borderRadius: '4px',
    border: '1px solid #ccc',
    width: '100%',
    maxWidth: '300px',
};

const buttonStyle: CSSProperties = {
    padding: '10px 20px',
    borderRadius: '4px',
    border: 'none',
    cursor: 'pointer',
};

export default Home;