export type AuthResponse = {
    accessToken: string;
    refreshToken: string;
}

export type AuthErrorResponse = {
    status: number;
    error: string;
    message: string;
}

export async function login(email: string, password: string) {
    try {
        const response = await fetch(`${process.env.BACKEND_URL}/api/auth/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        });
        if (response.ok) {
            const json = await response.json();
            return json as AuthResponse;
        } else {
            // Server returned an error
            const errorResponse = await response.json();
            return {status: response.status, error: errorResponse.error, message: "Failed to login"} as AuthErrorResponse;
        }
    } catch(error){
        console.error('Failed to login:', error);
        // throw new Error('Failed to login');
        return {error: error, message: "Failed to login"} as AuthErrorResponse;
    }
}

export async function register(email: string, password: string) {
    try {
        const response = await fetch(`${process.env.BACKEND_URL}/api/auth/register`, {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
        });
        if (response.ok) {
            const json = await response.json();
            return json as AuthResponse;
        }
        else {
        // Server returned an error
        const errorResponse = await response.json();
        return {error: errorResponse.error, message: "Failed to register"} as AuthErrorResponse;
    } 
    } catch(error){
        return {error: error, message: "Failed to register"} as AuthErrorResponse;
    }
}


export async function refresh(refreshToken: string) {
    console.log("refreshing token");
    try {
        const response = await fetch(`${process.env.BACKEND_URL}/api/auth/refresh`, {
            method: 'POST',
            headers: {
                RefreshToken: `${refreshToken}`
            }
        });
        console.log(response.status);
        if (response.ok) {
            const json = await response.json();
            return json as AuthResponse;
        } else {
            // Server returned an error
            const errorResponse = await response.json();
            throw new Error(errorResponse.message || 'Failed to login');
        }
    } catch (error) {
        // Handle any errors here
        console.error(error);
        return {error: error, message: "Failed to refresh token"} as AuthErrorResponse;
    }
}
