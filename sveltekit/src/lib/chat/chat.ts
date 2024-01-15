export type ChatCompletitionResponse = {
    response: string;
}

export type ChatCompletitionErrorResponse = {
    error: string;
    message: string;
}

export async function ChatCompletition(email: string, password: string) {
    try {
        const response = await fetch('http://localhost:8080/api/ChatCompletetionTest', {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
        });
        const json = await response.json();
        if (response.ok) {
            return json as ChatCompletitionResponse;
        }
    } catch(error){
        // Handle any errors here
        console.error(error);
        return { error: error, message: "Failed to login" } as unknown as ChatCompletitionResponse;
    }
}