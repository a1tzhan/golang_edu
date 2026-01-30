import { useState } from "react";

function RegisterPage() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");


    const handleSubmit = async (e: { preventDefault: () => void; }) => {
        e.preventDefault();
        const response = await fetch("http://localhost:8000/auth/register", {
            headers: {
                "Content-Type": "application/json",
            },
            method: "POST",
            body: JSON.stringify({
                email: email,
                password: password,
                confirmPassword: confirmPassword
            }),
        });

        console.log("Registration response:", response);

    };


    return (
        <div className="flex items-center justify-center h-screen">
            <div className="border border-gray-300 rounded-md w-[400px] p-6">
                <h1 className="text-gray-400 text-4xl text-center pb-10">Custom University Portal</h1>
                <h2 className="text-gray-600 text-3xl pb-4">Register Page </h2>
                <div className="">
                    <form className="flex flex-col gap-4" onSubmit={handleSubmit}>
                        <div>
                            <label htmlFor="email">Email:</label>
                            <input type="email" id="email" name="email" value={email} onChange={(e) => setEmail(e.target.value)} required className="border border-gray-300 rounded-md w-full h-10" />
                        </div>
                        <div>
                            <label htmlFor="password">Password:</label>
                            <input type="password" id="password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} required className="border border-gray-300 rounded-md w-full h-10" />
                        </div>
                        <div>
                            <label htmlFor="confirm-password">Confirm Password:</label>
                            <input type="password" id="confirm-password" name="confirm-password" value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} required className="border border-gray-300 rounded-md w-full h-10" />
                        </div>
                        <div>
                            <button className="bg-blue-500 text-white text-xl rounded-md w-full h-10" type="submit">Register</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    );
}

export { RegisterPage };
