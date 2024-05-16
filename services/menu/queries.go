package menu

const _getRestaurantMenu = `
SELECT meals.*, menus.*
FROM meals
JOIN menu_meals on meals.id = menu_meals.meals_id
JOIN menus on menus.id = menu_meals.menus.id = menus.id
JOIN restaurants on restaurants.id = menu.restaurant_id
JOIN branches on branches.restaurant_id = restaurants.id
JOIN tables on branches.id = tables.branches.id
where tables.id = ? and menus.active = true
group by(meals.tag);
`
