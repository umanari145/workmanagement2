const path = require('path')

module.exports = {
    configureWebpack: {
        resolve: {
          alias: {
            '@': path.join(__dirname, '/src') // 1. @の参照先の変更
          }
        }
    }
}
