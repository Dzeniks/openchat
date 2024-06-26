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
            try {
                const json = await response.json();
                if (json.error) {
                    return {status: response.status, error: json.error, message: "Failed to login"} as AuthErrorResponse;
                }
            }
            catch {
                return {status: response.status, error: response.statusText, message: "Failed to login"} as AuthErrorResponse;
            }
            return {status: response.status, error: response.statusText, message: "Failed to login"} as AuthErrorResponse;
        }
    } catch(error){
        return {status:500, "error": "500", message: "Failed to login"} as AuthErrorResponse;
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
        return {status: response.status, error: response.statusText, message: "Failed to register"} as AuthErrorResponse;
    } 
    } catch(error){
        return {error: error, message: "Failed to register"} as AuthErrorResponse;
    }
}


export async function refresh(refreshToken: string) {
    try {
        const response = await fetch(`${process.env.BACKEND_URL}/api/auth/refresh`, {
            method: 'POST',
            headers: {
                RefreshToken: `${refreshToken}`
            }
        });
        if (response.ok) {
            const json = await response.json();
            return json as AuthResponse;
        } else {
            return {status: response.status, error: response.statusText, message: "Failed to refresh token"} as AuthErrorResponse;
        }
    } catch (error) {
        return {error: error, message: "Failed to refresh token"} as AuthErrorResponse;
    }
}

export async function activate(refreshToken: string) {
    try {
        const response = await fetch(`${process.env.BACKEND_URL}/api/auth/activate`, {
            method: 'POST',
            headers: {
                RefreshToken: `${refreshToken}`
            }
        });
        if (response.ok) {
            const json = await response.json();
            return json as AuthResponse;
        } else {
            return {status: response.status, error: response.statusText, message: "Failed to activate"} as AuthErrorResponse;
        }
    } catch (error) {
        return {error: error, message: "Failed to activate"} as AuthErrorResponse;
    }
}

export async function auth(accessToken: string): Promise<boolean> {
    const response = await fetch(`${process.env.BACKEND_URL}/api/auth/`, {
        method: 'POST',
        headers: {
            Authorization: `${accessToken}`
        }
    });
    return response.ok;
}
