<link rel="stylesheet" href="/static/css/main.css">
<script src="/static/js/jquery.js"></script>
<script src="/static/js/moment.js"></script>
<script src="/static/js/livestamp.js"></script>

<header>
  <div class="heading">
    <span class="main-title">Nifty 50</span> 
    <span> - updated from NSE</span>
    <span class="livestamp" data-livestamp="{{.when}}"></span>
  </div>
</header>

<div class="tablesection-wrapper">
  <div class="tablesection gainers">

    <h2>Top 10 Gainers</h2>

    <table class="gridtable">
      <thead>
        <tr>
          <th>Symbol</th>
          <th>LTP</th>
          <th>% Change</th>
          <th>Traded Qty</th>
          <th>Value (in Lakhs)</th>
          <th>Open</th>
          <th>High</th>
          <th>Low</th>
          <th>Prev. Close</th>
          <th>Latest Ex Date</th>
          <th>CA</th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .grecords}}
        <tr>
          <td class="text">{{$record.Symbol}}</td>
          <td class="number">{{$record.Ltp}}</td>
          <td class="number">{{$record.Netprice}}</td>
          <td class="number">{{$record.TradedQuantity}}</td>
          <td class="number">{{$record.TurnoverInLakhs}}</td>
          <td class="number">{{$record.OpenPrice}}</td>
          <td class="number">{{$record.HighPrice}}</td>
          <td class="number">{{$record.LowPrice}}</td>
          <td class="number">{{$record.PreviousPrice}}</td>
          <td>{{$record.LastCorpAnnouncementDate}}</td>
          <td ><img style="" title="{{$record.LastCorpAnnouncement}}" src="/static/img/note_ico.gif"></td>
        </tr>
        {{end}}
      </tbody>
    </table>
  </div>

  <div class="tablesection losers">
    <h2>Top 10 Losers</h2>

    <table class="gridtable">
      <thead>
        <tr>
          <th>Symbol</th>
          <th>LTP</th>
          <th>% Change</th>
          <th>Traded Qty</th>
          <th>Value (in Lakhs)</th>
          <th>Open</th>
          <th>High</th>
          <th>Low</th>
          <th>Prev. Close</th>
          <th>Latest Ex Date</th>
          <th>CA</th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .lrecords}}
        <tr>
          <td class="text">{{$record.Symbol}}</td>
          <td class="number">{{$record.Ltp}}</td>
          <td class="number">{{$record.Netprice}}</td>
          <td class="number">{{$record.TradedQuantity}}</td>
          <td class="number">{{$record.TurnoverInLakhs}}</td>
          <td class="number">{{$record.OpenPrice}}</td>
          <td class="number">{{$record.HighPrice}}</td>
          <td class="number">{{$record.LowPrice}}</td>
          <td class="number">{{$record.PreviousPrice}}</td>
          <td>{{$record.LastCorpAnnouncementDate}}</td>
          <td ><img style="" title="{{$record.LastCorpAnnouncement}}" src="/static/img/note_ico.gif"></td>
        </tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>