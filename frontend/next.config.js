/** @type {import('next').NextConfig} */

/**
 * The type of build output.
 * - `undefined`: The default build output, `.next` directory, that works with production mode `next start` or a hosting provider like Vercel
 * - `'standalone'`: A standalone build output, `.next/standalone` directory, that only includes necessary files/dependencies. Useful for self-hosting in a Docker container.
 * - `'export'`: An exported build output, `out` directory, that only includes static HTML/CSS/JS. Useful for self-hosting without a Node.js server.
 */

module.exports = {

    output: "standalone"

};

