
// components/card/card.js
Component({
  properties: {
    'magazine': {
      type: String,
      value: '期刊名称'
    },
    'year': {
      type: Number,
      value: 2021
    },
    'issue': {
      type: Number,
      value: 0
    },
    'title': {
      type: String,
      value: '文章标题'
    },
    'authors': {
      type: [String],
      value: ['auhtor']
    },
    'paper_id': {
      type: String
    }
  },
  data: {

  },
  methods: {
    jump_detail(e) {
      wx.navigateTo({
        url: '/pages/detail/index?id='+e.currentTarget.dataset.id,
      })
    }
  }
})
