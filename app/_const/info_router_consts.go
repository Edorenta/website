package _const

const ( //strings?
	IR_PORT	uint16 = 20001	//info routing port
	N_BROKERS int = 3		//brokers: [pepperstone,forex.com,xm.com]
	//MT4 global data queries
	MSG_QUOTE_CONCAT = "001"
	MSG_ACCOUNT_INFO_CONCAT = "002"
	MSG_MARKET_INFO_CONCAT = "003"
	//MT4 specific account queries
	MSG_ACCOUNT_NUMBER = "101"
	MSG_ACCOUNT_NAME = "102"
	MSG_ACCOUNT_BROKER_NAME = "103"
	MSG_ACCOUNT_SERVER = "104"
	MSG_ACCOUNT_CURRENCY = "105"
	MSG_ACCOUNT_CREDIT = "106"
	MSG_ACCOUNT_LEVERAGE = "107"
	MSG_ACCOUNT_STOPOUT_LVL = "108"
	MSG_ACCOUNT_STOPOUT_MODE = "109"
	MSG_ACCOUNT_MARGIN = "110"
	MSG_ACCOUNT_FREE_MARGIN = "111"
	MSG_ACCOUNT_BALANCE = "112"
	MSG_ACCOUNT_EQUITY = "113"
	MSG_ACCOUNT_PROFIT = "114"
	MSG_ACCOUNT_RELATIVE_DRAWDAWN = "115"
	MSG_ACCOUNT_ABSOLUTE_DRAWDAWN = "116"
	//MT4 specific quote queries
	MSG_TRADEALLOWED = "201"
	MSG_LEVERAGE = "202"
	MSG_RECQUIRED_MARGIN = "203"
	MSG_LOTSIZE = "204"
	MSG_MIN_LOT = "205"
	MSG_MAX_LOT = "206"
	MSG_LOT_STEP = "208"
	MSG_TICK_SIZE = "209"
	MSG_TICK_VALUE = "210"
	MSG_LOT_VALUE = "211"
	MSG_STOP_LEVEL = "212"
	MSG_FREEZE_LEVEL = "213"
	MSG_BID = "214"
	MSG_ASK = "215"
	MSG_SPREAD = "216"
	MSG_SWAP_LONG = "217"
	MSG_SWAP_SHORT = "218"
	MSG_STARTING = "219"
	MSG_EXPIRY = "220"
	MSG_STREAK = "221"
	MSG_WINRATE = "222"
	MSG_ABSOLUTE_DRAWDOWN = "223"
	MSG_RELATIVE_DRAWDOWN = "224"
	MSG_LOSS_SUM = "225"
	MSG_PROFIT_SUM = "226"
	MSG_PNL = "227"
	MSG_ORDER_SUM = "228"
	//WS ReqRes queries
	MSG_START_TERM_FEED = "301"
	MSG_STOP_TERM_FEED = "302"
	MSG_START_WEB_FEED = "303"
	MSG_STOP_WEB_FEED = "304"
	MSG_START_LOG = "305"
	MSG_STOP_LOG = "306"
)
