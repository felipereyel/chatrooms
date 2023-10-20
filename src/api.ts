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
    })

    if (!r.ok) {
        throw new Error(await r.text())
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
    })

    if (!r.ok) {
        throw new Error(await r.text())
    }
}

export const cookielogout = async () => {
    document.cookie = `${cookieName}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`
}

export const isloggedin = () => {
    return document.cookie.includes(cookieName)
};

export const listrooms = async () => {
    const r = await fetch(`/_api/rooms/`)

    if (!r.ok) {
        throw new Error(await r.text())
    }

    return await r.json()
};

export const getroom = async (id: string) => {
    const r = await fetch(`/_api/rooms/${id}`)

    if (!r.ok) {
        throw new Error(await r.text())
    }

    return await r.json()
};

export const createroom = async (name: string) => {
    const r = await fetch(`/_api/rooms`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ name })
    })

    if (!r.ok) {
        throw new Error(await r.text())
    }

    return await r.json()
};

export const listposts = async (roomid: string) => {
    const r = await fetch(`/_api/rooms/${roomid}/posts/`)

    if (!r.ok) {
        throw new Error(await r.text())
    }

    return await r.json()
};

export const createpost = async (roomid: string, content: string) => {
    const r = await fetch(`/_api/rooms/${roomid}/posts`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ content })
    })

    if (!r.ok) {
        throw new Error(await r.text())
    }

    return await r.json()
};

const _initRoom = async (roomid: string) => {
    const hundredPosts = Array.from({ length: 100 }, (_, i) => i + 1).map(i => ({
        content: `Post ${i}`,
    }))

    for (const post of hundredPosts) {
        await createpost(roomid, post.content)
    }
}