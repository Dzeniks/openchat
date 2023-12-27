export type AuthResponse = {
    accessToken: string;
    refreshToken: string;
}

export async function login(email: string, password: string) {
    try {
        const response = await fetch('http://localhost:8080/api/auth/login', {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
        });
        const json = await response.json();
        if (response.ok) {
            return json as AuthResponse;
        }
    } catch(error){
        // Handle any errors here
        console.error(error);
        
        // Return an error message or throw an exception
        throw new Error('Failed to login');
    }
}

export async function register(email: string, password: string) {
    try {
        const response = await fetch('http://localhost:8080/api/auth/register', {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
        });
        const json = await response.json();
        if (response.ok) {
            return json as AuthResponse;
        }
    } catch(error){
        // Handle any errors here
        console.error(error);
        
        // Return an error message or throw an exception
        throw new Error('Failed to register');
    }
}


export async function refresh(refreshToken: string) {
    console.log("refreshing token");
    try {
        const response = await fetch('http://localhost:8080/api/auth/refresh', {
            method: 'POST',
            headers: {
                RefreshToken: `${refreshToken}`
            }
        });
        console.log(response.status);
        if (response.ok) {
            const json = await response.json();
            return json as AuthResponse;
        } if (response.status === 400) {
            const json = await response.json();
            console.log(json);
        }
    } catch (error) {
        // Handle any errors here
        console.error(error);
        
        // Return an error message or throw an exception
        throw new Error('Failed to refresh token');
    }
}