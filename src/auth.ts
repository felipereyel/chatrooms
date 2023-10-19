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