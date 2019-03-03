grammar Path;

query
  : AND? andExpr AND?
  | AND? orExpr AND?
  | AND? tag AND?
  | AND? anyExpr AND?
  | AND? noneExpr AND?
  |
  ;

noneExpr
  : NONE
  ;

anyExpr
  : ANY
  | ANY (AND ANY)+
  ;

andExpr
  : orExpr (AND orExpr)+
  | orExpr (AND tag)+
  | tag (AND orExpr)+
  | tag (AND tag)+
  ;

orExpr
  : tag (OR tag)+
  ;

tag
  : classifierTagExpr
  | nameTagExpr
  ;

classifierTagExpr
  : classifier=STRING op=(EQ | NEQ) name=ANY
  ;

nameTagExpr
  : name=STRING
  | op=NOT name=STRING
  | classifier=STRING op=(EQ | NEQ) name=STRING
  ;

AND : '/' ;
ANY : ('*' | '**' ) ;
EQ : '=' ;
NEQ : '!=' ;
NONE : '!*' ;
NOT : '!' ;
OR : ',' ;
STRING : ~('/' | ',' | '=' | '!')+ ;
