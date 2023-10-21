import jwt_decode from "jwt-decode";

const cookieName = "chatrooms:jwt";

export const apilogin = async (username: string, password: string) => {
    const r = await fetch(`/_api/users/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            username, password
        })
    });

    if (!r.ok) {
        throw new Error(await r.text());
    }
}

export const apiregister = async (username: string, password: string) => {
    const r = await fetch(`/_api/users/register`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            username, password
        })
    });

    if (!r.ok) {
        throw new Error(await r.text());
    }
}

export const cookielogout = async () => {
    document.cookie = `${cookieName}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
}

export const isloggedin = () => {
    return document.cookie.includes(cookieName);
};

export const listrooms = async () => {
    const r = await fetch(`/_api/rooms/`);

    if (!r.ok) {
        throw new Error(await r.text());
    }

    return await r.json();
};

export const getroom = async (id: string) => {
    const r = await fetch(`/_api/rooms/${id}`);

    if (!r.ok) {
        throw new Error(await r.text());
    }

    return await r.json();
};

export const createroom = async (name: string) => {
    const r = await fetch(`/_api/rooms/`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ name })
    });

    if (!r.ok) {
        throw new Error(await r.text());
    }

    return await r.json();
};

export const listposts = async (roomid: string) => {
    const r = await fetch(`/_api/rooms/${roomid}/posts`);

    if (!r.ok) {
        throw new Error(await r.text());
    }

    return await r.json();
};

export const createpost = async (roomid: string, content: string) => {
    const r = await fetch(`/_api/rooms/${roomid}/posts`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ content })
    });

    if (!r.ok) {
        throw new Error(await r.text());
    }
};

export const getRoomWs = (roomid: string) => {
    return new WebSocket(`ws://${window.location.host}/_api/rooms/${roomid}/ws`);
}

export const getUsername = (): string => {
    try {
        const cookie = document.cookie.split(";").find(c => c.includes(cookieName));
        if (!cookie) {
            return "Anon";
        }
    
        const token = cookie.split("=")[1];
    
        const decoded: any = jwt_decode(token);
        return decoded.sub ?? "Anon";
    } catch (e) {
        return "Anon";
    }
}