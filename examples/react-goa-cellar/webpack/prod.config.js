// Webpack config for creating the production bundle.
require('babel-core/polyfill');
var path = require('path');
var fs = require('fs');
var webpack = require('webpack');
var CleanPlugin = require('clean-webpack-plugin');
var ExtractTextPlugin = require('extract-text-webpack-plugin');
var strip = require('strip-loader');

var appPath = path.join(__dirname, '/..');
var assetsPath = path.join(appPath, '/dist');

var babelrc = fs.readFileSync(path.join(appPath, '/.babelrc'));
var babelrcObject = {};

var isRehearsal = (process.env.NODE_ENV == 'rehearsal');
console.log("NODE ENV => " + process.env.NODE_ENV);
console.log("isRehearsal => " + isRehearsal);

try {
  babelrcObject = JSON.parse(babelrc);
} catch (err) {
  console.error('==>     ERROR: Error parsing your .babelrc.');
  console.error(err);
}

var babelrcObjectEnv = babelrcObject.env && (babelrcObject.env.production || {});
var babelLoaderQuery = Object.assign({}, babelrcObject, babelrcObjectEnv);
delete babelLoaderQuery.env;

module.exports = {
  devtool: 'cheap-module-source-map',
  context: appPath,
  entry: {
    'main': [
      path.join(appPath, 'src/index.js')
    ]
  },
  frameworks: ["webpack"],
  output: {
    path: assetsPath,
    filename: "index.js"
  },
  module: {
    loaders: [
      {test: /\.(jpe?g|png|gif|svg)$/, loader: 'url', query: {limit: 40960}},
      {
        test: /\.scss$/,
        loader: 'style!css?modules&importLoaders=2&sourceMap&localIdentName=[local]___[hash:base64:5]!autoprefixer?browsers=last 2 version!sass?outputStyle=expanded&sourceMap'
      },
      {test: /\.css$/, loader: "style-loader!css-loader"},
      {test: /\.less$/, loader: "style-loader!css-loader!less-loader"},
      {test: /\.js$/, exclude: /node_modules/, loaders: ['babel?' + JSON.stringify(babelLoaderQuery), 'eslint-loader']},
      {test: /\.json$/, loader: 'json-loader'},
      {test: /\.md/, loaders: ["html-loader", "markdown-loader"]},
      {test: /\.gif$/, loader: "url-loader?mimetype=image/png"},
      {test: /\.woff(2)?(\?v=[0-9].[0-9].[0-9])?$/, loader: "url-loader?mimetype=application/font-woff"},
      {test: /\.(ttf|eot|svg)(\?v=[0-9].[0-9].[0-9])?$/, loader: "file-loader?name=[name].[ext]"}
    ]
  },
  progress: true,
  resolveLoader: {
    root: path.join(appPath, 'node_modules'),
  },
  plugins: [
    new CleanPlugin([assetsPath]),

    // css files from the extract-text-plugin loader
    new ExtractTextPlugin('bundle.css', {allChunks: true}),

    // ignore dev config
    new webpack.IgnorePlugin(/\.\/dev/, /\/config$/),

    // set global vars
    new webpack.DefinePlugin({
      __DEVELOPMENT__: false,
      __DEVTOOLS__: false,
      'process.env': {
        // Useful to reduce the size of client-side libraries, e.g. react
        NODE_ENV: JSON.stringify('production')
      }
    }),

    // optimizations
    new webpack.optimize.DedupePlugin(),
    new webpack.optimize.OccurenceOrderPlugin(),
    new webpack.optimize.UglifyJsPlugin({
      compress: {
        warnings: false
      }
    }),
  ]
};
