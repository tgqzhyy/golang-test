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