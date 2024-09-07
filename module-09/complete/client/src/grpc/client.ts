import {createPromiseClient, PromiseClient} from "@connectrpc/connect";
import {createConnectTransport} from "@connectrpc/connect-web";
import {ServiceType} from "@bufbuild/protobuf";
import {useMemo} from "react";

const transport = createConnectTransport({
    baseUrl: "http://localhost:50052",
    // Not needed. Web browsers use HTTP/2 automatically.
    // httpVersion: "1.1"
});

export function useClient<T extends ServiceType>(service: T): PromiseClient<T> {
    // We memoize the client, so that we only create one instance per service.
    return useMemo(() => createPromiseClient(service, transport), [service]);
}