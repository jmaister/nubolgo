const path = require('path');

const HtmlWebPackPlugin = require("html-webpack-plugin");
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

const PATHS = {
    src: path.join(__dirname, 'src/'),
    build: path.join(__dirname, '../dist/'),
    public: path.join(__dirname, 'public/'),
};

const htmlPlugin = new HtmlWebPackPlugin({
    template: path.join(__dirname, '/src/index.html'),
    filename: "./index.html",
    favicon: PATHS.public + "favicon.ico"
});

module.exports = {
    entry: PATHS.src + 'index.js',
    output: {
        path: PATHS.build,
        filename: 'index.js'
    },
    optimization: {
        // true for tree shaking
        usedExports: true
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader"
                }
            },
            {
                test: /\.scss$/,
                exclude: /node_modules/,
                use: [
                    "style-loader", // creates style nodes from JS strings
                    'css-loader', // translates CSS into CommonJS
                    "sass-loader" // compiles Sass to CSS, using Node Sass by default
                ]
            },
            {
                test: /\.ico$/,
                exclude: /node_modules/,
                use: {
                    loader: "image-loader"
                }
            }
        ]
    },
    plugins: [
        new CleanWebpackPlugin(),
        htmlPlugin,
        new CopyWebpackPlugin([
            { from: PATHS.public }
        ])
    ]
};
