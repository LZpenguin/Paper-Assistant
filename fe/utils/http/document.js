function document_20(token,callback) {
  wx.request({
    url: 'https://api.hust.online/doc-aid-test/document',
    success: function(res) {
      if(callback) {
        callback(res.data.data)
      }
    },
    header: {
      'Authorization': 'Bearer '+ token
    }
  })
}

module.exports = {document_20}