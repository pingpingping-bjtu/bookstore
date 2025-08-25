package service

import "bookstore-manager/repository"

type AdminDashboardService struct {
	AdminDashboardDAO *repository.AdminDashboardDAO
}

// DashboardStats 仪表盘统计数据
type DashboardStats struct {
	TotalBooks   int64   `json:"total_books"`
	TotalOrders  int64   `json:"total_orders"`
	TotalUsers   int64   `json:"total_users"`
	TotalRevenue float64 `json:"total_revenue"`
	RecentBooks  []Book  `json:"recent_books"`
}

// Book 简化的图书信息
type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Price     int    `json:"price"`
	CoverURL  string `json:"cover_url"`
	CreatedAt string `json:"created_at"`
}

func (d *AdminDashboardService) GetDashboardStats() (DashboardStats, error) {
	//获取图书总数
	var stats DashboardStats
	totalBooks, err := d.AdminDashboardDAO.GetTotalBooks()
	if err != nil {
		return stats, err
	}
	//获取订单总数
	totalOrders, err := d.AdminDashboardDAO.GetTotalOrders()
	if err != nil {
		return stats, err
	}
	//获取用户总数
	totalUsers, err := d.AdminDashboardDAO.GetTotalUsers()
	if err != nil {
		return stats, err
	}
	//获取总收入
	totalRevenue, err := d.AdminDashboardDAO.GetTotalRevenue()
	if err != nil {
		return stats, err
	}
	//获取最近添加的书籍
	recentBooks, err := d.AdminDashboardDAO.GetRecentBooks()
	if err != nil {
		return stats, err
	}
	stats.TotalBooks = totalBooks
	stats.TotalOrders = totalOrders
	stats.TotalUsers = totalUsers
	stats.TotalRevenue = float64(totalRevenue)
	for _, book := range recentBooks {
		stats.RecentBooks = append(stats.RecentBooks, Book{
			ID:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Price:     book.Price,
			CoverURL:  book.CoverURL,
			CreatedAt: book.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return stats, nil

}

func NewAdminDashboardService() *AdminDashboardService {
	return &AdminDashboardService{AdminDashboardDAO: repository.NewAdminDashboardDAO()}
}
