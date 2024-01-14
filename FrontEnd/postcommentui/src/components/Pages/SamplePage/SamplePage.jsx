import React, { useEffect, useState } from "react";

const SamplePage = () => {
    const [count, setCount] = useState(0);

    useEffect(() => {
        const interval = setInterval(() => {
            setCount(count => count + 1);
        }, 1000)

        return () => {
            clearInterval(interval);
        }
    }, [])

    return (
        <div>
            <h1>Sample Page, {count}s</h1>
        </div>
    );
}

export default SamplePage