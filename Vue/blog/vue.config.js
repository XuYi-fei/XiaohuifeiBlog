module.exports = {
    publicPath: "/",

    devServer: {
        host: '0.0.0.0',
        port: 80,
        proxy: {
            '/api': {
                // target: '127.0.0.1:80', // 接口域名
                target: '121.89.200.190:80', // 接口域名
                changeOrigin: true, // 是否跨域
                ws: false, // 是否代理 websockets
                secure: false, // 是否https接口
                pathRewrite: { // 路径重置
                    '^/api': ''
                }
            }
        }
    },
    productionSourceMap: false,
    // css: {
    //     // 将组件内部的css提取到一个单独的css文件（只用在生产环境）
    //
    //     // 也可以是传递给 extract-text-webpack-plugin 的选项对象
    //
    //     extract: true, // 允许生成 CSS source maps?
    //
    //     sourceMap: false, // pass custom options to pre-processor loaders. e.g. to pass options to // sass-loader, use { sass: { ... } }
    //
    //     loaderOptions: {}, // Enable CSS modules for all css / pre-processor files. // This option does not affect *.vue files.
    //
    //     requireModuleExtension: false
    // },
    // productionGzip: true,
    // productionGzipExtensions: ['js', 'css'],
    // test: /\.md$/,
    //
    // cnpm install highlight.js -S
}
