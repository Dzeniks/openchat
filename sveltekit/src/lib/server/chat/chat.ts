

const postData = async (data: any) => {
    try {
        const response = await fetch("http://localhost:8080/api/chat/ChatCompletitions", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (!response.ok) {
            throw new Error('Failed to POST data');
        }

        const result = await response.json();
        return result;
    } catch (error) {
        console.error(error);
        throw error;
    }
};
