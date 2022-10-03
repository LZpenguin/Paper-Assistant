function get_token(context,callback) {
  wx.getStorage({
    key: 'token',
    success: res => {
      context.setData({
        'token': res.data
      })
      callback(res.data)
    }
  })
}

module.exports = {get_token}