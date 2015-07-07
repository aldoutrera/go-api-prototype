;(function($, window, document) {

  var TourModel = Backbone.Model.extend({
    defaults: {
      id: null,
      name: null,
    }
  });

  var ToursCollection = Backbone.Collection.extend({
    url: '/tours',
    model: TourModel,
    parse: function(data) {
      return data.tours;
    }
  });

  var ToursListItemView = Backbone.View.extend({
    tagName: 'li',
    className: 'tour',
    render: function() {
      // var html = '<b>Id: </b>' + this.model.get('id');
      // html += ', ' + this.model.get('name');
      var html = this.template(this.model.toJSON());
      this.$el.html(html);
      return this;
    },
    initialize: function() {
      this.template = _.template($('#tour-item-tmpl').html());
    }
  });

  var ToursListView = Backbone.View.extend({
    el: '#tours-list',
    initialize: function() {
      this.listenTo(this.collection, 'sync change', this.render);
      this.collection.fetch();
      this.render();
    },
    render: function() {
      this.collection.each(function(model) {
        var item = new ToursListItemView({ model: model });
        $(this.$el.selector).append(item.render().$el);
      }, this);
      return this;
    }
  });

  var tours = new ToursCollection();
  var toursView = new ToursListView({ collection: tours });
  // tours.fetch();

})(jQuery, window, document);
