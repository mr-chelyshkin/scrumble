package torrent_fetcher

type App struct {

}

func (a App) Name() string {
	return "torrent-fetcher"
}

func (a App) ThirdParty(e chan <- error) {
	return
}

func (a App) Run() error {
	return nil
}