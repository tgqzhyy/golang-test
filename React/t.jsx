function Comment(props) {
    return(
        <div className="Comment">
            <div className="UserInfo">
                <img className="Avater"
                src={props.author.avatarUrl}
                 alt={props.author.name}
                />
                <div className="UserInfo-name">
                    {props.author.name}
                </div>
            </div>
            <div className="Comment-text">
                {props.text}
            </div>
            <div className="Comment-date">
                {formatDate(props.date)}
            </div>
        </div>
    );
}


function Avatar(props) {
    return(
        <img className="Avatar"
            src={props.user.avatarUrl}
            alt={props.user.name}
        />
    );
}


function Comment(props) {
    return(
        <div className="Comment">
            <div className="UserInfo">
                <Avatar user={props.author} />
                <div className="UserInfo-name">
                    {props.author.name}
                </div>
            </div>
            <div className="Comment-text">
                {props.text}
            </div>
            <div className="Comment-date">
                {formatDate(props.date)}
            </div>
        </div>
    );
}


function sum(a,b) {
    return a+b;
}


function withdraw(account,amount) {
    account.total -=amount;
}

/**
 *
 */

function WarningBanner(props) {
    if (!props.warn){
        return null;
    }

    return (
        <div className="warning">
            警告！
        </div>
    );
}

class Page extends React.Component{
    constructor(props) {
        super(props);
        this.state={showWarning:true}
        this.handleToggleClick =this.handleToggleClick.bind(this)
    }

    handleToggleClick(){
        this.setState(prevState=>(
            {
                showWarning: !prevState.showWarning
            }
        ));
    }

    render(){
        return(
            <div>
                <WarningBanner warn={this.state.showWarning}></WarningBanner>
                <button onClick={this.handleToggleClick}>
                    {this.state.showWarning?'隐藏':'显示'}
                </button>
            </div>
        );
    }
}

ReactDom.render(
    <Page></Page>,
    document.getElementById('example')
)