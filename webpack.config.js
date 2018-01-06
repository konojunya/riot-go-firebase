const path = require("path");
const webpack = require("webpack");

module.exports = {
  resolve: {
    extensions: [".js", ".tag"]
  },
  entry: __dirname + "/src/app.js",
  output: {
    path: __dirname + "/public/js/bundle",
    filename: "bundle.js",
    libraryTarget: "umd"
  },
  module: {
    rules: [
      {
        test: /\.tag$/,
        enforce: "pre",
        exclude: /node_modules/,
        use: [
          {
            loader: "riotjs-loader",
            options: {
              debug: true
            }
          }
        ]
      },
      {
        test: /\.js$|\.tag$/,
        enforce: "post",
        exclude: /node_modules/,
        use: [
          {
            loader: "babel-loader",
            options: {
              presets: ["es2015-riot"]
            }
          }
        ]
      }
    ]
  },
  plugins: [
    new webpack.ProvidePlugin({
      riot: "riot"
    })
  ]
}