package gbatis

type SqlSession struct {
}

func (self *SqlSession) SelectOne(sqlId string, parameter interface{}) (interface{}, error) {
    return nil
}

func (self *SqlSession) SelectList(sqlId string, parameter interface{}) ([]interface{}, error) {
    return nil
}

func (self *SqlSession) Execute(sqlId string, parameter interface{}) error {
    
}

func (self *SqlSession) ExecuteMany(sqlId string, parameter []interface{}) error {

}
